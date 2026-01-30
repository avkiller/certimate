#!/usr/bin/env python3
import logging
import json
import mimetypes
import tempfile
import os
import random
import re
import shutil
import time
from urllib import request
from urllib.error import HTTPError

GITHUB_REPO = "certimate-go/certimate"
GITEE_REPO = "certimate-go/certimate"
GITEE_TOKEN = os.getenv("GITEE_TOKEN", "")

SYNC_MARKER = "SYNCING FROM GITHUB, PLEASE WAIT ..."
TEMP_DIR = tempfile.mkdtemp()

logging.basicConfig(level=logging.INFO)


def do_httpreq(url, method="GET", headers=None, data=None):
    req = request.Request(url, data=data, method=method)
    headers = headers or {}
    for key, value in headers.items():
        req.add_header(key, value)

    try:
        with request.urlopen(req) as resp:
            resp_data = resp.read().decode("utf-8")
            if resp_data:
                try:
                    return json.loads(resp_data)
                except json.JSONDecodeError:
                    pass
            return None
    except HTTPError as e:
        errmsg = ""
        if e.readable():
            try:
                errmsg = e.read().decode('utf-8')
                errmsg = errmsg.replace("\r", "\\r").replace("\n", "\\n")
            except:
                pass
        logging.error(f"Error occurred when sending request: status={e.status}, response={errmsg}")
        raise e
    except Exception as e:
        raise e


def get_github_stable_release():
    page = 1
    while True:
        releases = do_httpreq(
            url=f"https://api.github.com/repos/{GITHUB_REPO}/releases?page={page}&per_page=100",
            headers={"Accept": "application/vnd.github+json"},
        )
        if not releases or len(releases) == 0:
            break

        for release in releases:
            release_name = release.get("name", "")
            if re.match(r"^v[0-9]", release_name):
                if any(
                    x in release_name
                    for x in ["alpha", "beta", "rc", "preview", "test", "unstable"]
                ):
                    continue
                return release

        page += 1

    return None


def get_gitee_release_list():
    page = 1
    list = []
    while True:
        releases = do_httpreq(
            url=f"https://gitee.com/api/v5/repos/{GITEE_REPO}/releases?access_token={GITEE_TOKEN}&page={page}&per_page=100",
        )
        if not releases or len(releases) == 0:
            break

        list.extend(releases)
        page += 1

    return list


def get_gitee_release_by_tag(tag_name):
    releases = get_gitee_release_list()
    for release in releases:
        if release.get("tag_name") == tag_name:
            return release

    return None


def delete_gitee_release(release_info):
    if not release_info:
        raise ValueError("Release info is invalid")

    release_id = release_info.get("id", "")
    release_name = release_info.get("tag_name", "")
    if not release_id:
        raise ValueError("Release ID is missing")

    attachpage = 1
    attachfiles = []
    while True:
        releases = do_httpreq(
            url=f"https://gitee.com/api/v5/repos/{GITEE_REPO}/releases/{release_id}/attach_files?access_token={GITEE_TOKEN}&page={attachpage}&per_page=100",
        )
        if not releases or len(releases) == 0:
            break

        attachfiles.extend(releases)
        attachpage += 1

    for attachfile in attachfiles:
        attachfile_id = attachfile.get("id")
        attachfile_name = attachfile.get("name")
        logging.info("Trying to delete Gitee attach file: %s/%s", release_name, attachfile_name)
        do_httpreq(
            url=f"https://gitee.com/api/v5/repos/{GITEE_REPO}/releases/{release_id}/attach_files/{attachfile_id}?access_token={GITEE_TOKEN}",
            method="DELETE",
        )

    logging.info("Trying to delete Gitee release: %s", release_name)
    do_httpreq(
        url=f"https://gitee.com/api/v5/repos/{GITEE_REPO}/releases/{release_id}?access_token={GITEE_TOKEN}",
        method="DELETE",
    )


def create_gitee_release(name, tag, body, prerelease, gh_assets):
    release_info = do_httpreq(
        f"https://gitee.com/api/v5/repos/{GITEE_REPO}/releases?access_token={GITEE_TOKEN}",
        method="POST",
        headers={"Content-Type": "application/json"},
        data=json.dumps({
            "tag_name": tag,
            "name": name,
            "body": SYNC_MARKER,
            "prerelease": prerelease,
            "target_commitish": "",
        }).encode("utf-8"),
    )

    if not release_info or "id" not in release_info:
        return None
    logging.info("Gitee release created")

    release_id = release_info["id"]

    assets_dir = os.path.join(TEMP_DIR, "assets")
    os.makedirs(assets_dir, exist_ok=True)

    gh_assets = gh_assets or []
    for asset in gh_assets:
        logging.info("Tring to download asset from GitHub: %s", asset["name"])

        opener = request.build_opener()
        request.install_opener(opener)
        download_ts = time.time()
        download_url = asset.get("browser_download_url")
        download_path = os.path.join(assets_dir, asset["name"])
        def _hook(blocknum, blocksize, totalsize):
            nonlocal download_ts
            TIMESPAN = 5 # print progress every 5sec
            ts = time.time()
            pct = min(round(100 * blocknum * blocksize / totalsize, 2), 100)
            if (ts - download_ts < TIMESPAN) and (pct < 100):
                return
            download_ts = ts
            logging.info(f"Downloading {download_url} >>> {pct}%")

        request.urlretrieve(download_url, download_path, _hook)

    for asset in gh_assets:
        logging.info("Tring to upload asset to Gitee: %s", asset["name"])

        boundary = '----boundary' + ''.join(random.choice('0123456789abcdef') for _ in range(16))
        print(f"Using boundary: {boundary}")
        with open(os.path.join(assets_dir, asset["name"]), 'rb') as f:
            attachfile_mime = mimetypes.guess_type(asset["name"])[0] or 'application/octet-stream'
            attachfile_req = []
            attachfile_req.append(f"--{boundary}")
            attachfile_req.append(f'Content-Disposition: form-data; name="file"; filename="{asset["name"]}"')
            attachfile_req.append(f"Content-Type: {attachfile_mime}")
            attachfile_req.append("")
            attachfile_req.append(f.read().decode('latin-1'))
            attachfile_req.append(f"--{boundary}--")
            attachfile_req.append("")
            attachfile_req = "\r\n".join(attachfile_req).encode('latin-1')

            do_httpreq(
                f"https://gitee.com/api/v5/repos/{GITEE_REPO}/releases/{release_id}/attach_files?access_token={GITEE_TOKEN}",
                method="POST",
                headers={'Content-Type': f'multipart/form-data; boundary={boundary}'},
                data=attachfile_req,
            )
            logging.info("Asset uploaded: %s", asset["name"])

    release_info = do_httpreq(
        f"https://gitee.com/api/v5/repos/{GITEE_REPO}/releases/{release_id}?access_token={GITEE_TOKEN}",
        method="PATCH",
        headers={"Content-Type": "application/json"},
        data=json.dumps({
            "tag_name": tag,
            "name": name,
            "body": f"**此发行版同步自 GitHub，完整变更日志请访问 https://github.com/{GITHUB_REPO}/releases/{tag} 查看。**\n\n**因 Gitee 存储空间容量有限，仅能保留最新一个发行版，如需其余版本请访问 GitHub 获取。**\n\n---\n\n" + body,
            "prerelease": prerelease,
        }).encode("utf-8"),
    )
    logging.info("Gitee release updated")
    return release_info


def main():
    try:
        # 获取 GitHub 最新稳定发行版
        github_release = get_github_stable_release()
        if not github_release:
            logging.warning("GitHub stable release not found. Foget to release?")
            return
        else:
            logging.info("GitHub stable release found: %s", github_release.get('name'))

        # 提取稳定版的信息
        release_name = github_release.get("name")
        release_tag = github_release.get("tag_name")
        release_body = github_release.get("body")
        release_prerelease = github_release.get("prerelease", False)
        release_assets = github_release.get("assets", [])

        # 检查 Gitee 是否已有同名发行版
        gitee_release = get_gitee_release_by_tag(release_tag)
        if gitee_release and gitee_release.get("body") == SYNC_MARKER:
            logging.warning("Gitee syncing release found, cleaning up...")
            delete_gitee_release(gitee_release)
        elif gitee_release:
            logging.info("Gitee release already exists, exit.")
            return

        # 同步发行版
        gitee_release = create_gitee_release(release_name, release_tag, release_body, release_prerelease, release_assets)
        if not gitee_release:
            logging.warning("Failed to create Gitee release.")
            return

        # 清除历史发行版
        gitee_release_list = get_gitee_release_list()
        for release in gitee_release_list:
            if release.get("tag_name") == release_tag:
                continue
            else:
                delete_gitee_release(release)

        logging.info("Sync release completed.")

    except Exception as e:
        logging.fatal(str(e))
        exit(1)

    finally:
        if os.path.exists(TEMP_DIR):
            shutil.rmtree(TEMP_DIR)


if __name__ == "__main__":
    main()
