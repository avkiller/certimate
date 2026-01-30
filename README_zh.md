<h1 align="center">🔒 Certimate</h1>

<div align="center">

[![Stars](https://img.shields.io/github/stars/certimate-go/certimate?style=flat)](https://github.com/certimate-go/certimate)
[![Forks](https://img.shields.io/github/forks/certimate-go/certimate?style=flat)](https://github.com/certimate-go/certimate)
[![Docker Pulls](https://img.shields.io/docker/pulls/certimate/certimate?style=flat)](https://hub.docker.com/r/certimate/certimate)
[![Release](https://img.shields.io/github/v/release/certimate-go/certimate?style=flat&sort=semver)](https://github.com/certimate-go/certimate/releases)
[![License](https://img.shields.io/github/license/certimate-go/certimate?style=flat)](https://mit-license.org/)
[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/certimate-go/certimate)

</div>

<div align="center">

[English](README.md) ｜ 简体中文

</div>

---

## 🚩 项目简介

完全开源免费的自托管 SSL 证书 ACME 工具，申请、部署、续期全流程自动化可视化，支持各大主流云厂商。

- **自托管**：私有化部署，所有数据本地化存储，掌控数据的隐私与安全。
- **零依赖**：无需安装数据库、运行时或复杂框架，一键启动，开箱即用。
- **低占用**：超轻量的资源开销，仅需 ~16 MB 内存，甚至可以运行在家用路由器。
- **易操作**：图形化界面，通过简单配置即可完成证书申请、部署和续期的自动化工作。

## 💡 功能特性

- 灵活的工作流编排方式，证书从申请到部署完全自动化。
- 支持申请单/多/泛域名证书、IP 地址证书，可选 RSA、ECC 私钥算法。
- 支持 DNS-01（即基于域名解析验证）、HTTP-01（即基于文件验证）两种质询方式。
- 支持 PEM、PFX、JKS 等多种格式输出证书。
- 支持 60+ 域名托管商（如阿里云、腾讯云、AWS、Cloudflare、GoDaddy 等，[点此查看完整清单](https://docs.certimate.me/zh-CN/docs/reference/providers#supported-dns-providers)）。
- 支持 110+ 部署目标（如 Kubernetes、CDN、WAF、负载均衡等，[点此查看完整清单](https://docs.certimate.me/zh-CN/docs/reference/providers#supported-hosting-providers)）。
- 支持邮件、钉钉、飞书、企业微信、Discord、Slack、Telegram 等多种通知渠道。
- 支持 Let's Encrypt、Actalis、Google Trust Services、SSL.com、ZeroSSL 等多种 ACME 证书颁发机构。
- 更多特性等待探索。

## 🚀 快速启动

**1 分钟运行 Certimate！**

<details>
<summary>👉 二进制安装：</summary>

从 [GitHub Releases](https://github.com/certimate-go/certimate/releases) 页面下载预先编译好的可执行文件压缩包，解压缩后在终端中执行：

```bash
./certimate serve
```

</details>

<details>
<summary>👉 Docker 安装：</summary>

```bash
docker run -d \
  --name certimate \
  --restart unless-stopped \
  -p 8090:8090 \
  -v /etc/localtime:/etc/localtime:ro \
  -v /etc/timezone:/etc/timezone:ro \
  -v $(pwd)/data:/app/pb_data \
  certimate/certimate:latest
```

</details>

浏览器中访问 `http://127.0.0.1:8090`。

初始的管理员账号及密码：

- 账号：`admin@certimate.fun`
- 密码：`1234567890`

即刻使用 Certimate。或者阅读文档中的其他内容以了解更多。

## 📄 使用手册

请访问文档站 [docs.certimate.me](https://docs.certimate.me/) 以阅读使用手册。

> （由于众所周知的原因，中国大陆用户可能需要 🪄 上网才能访问文档站。）

相关文章：

> - [《升级指南：迁移到 v0.4》](https://docs.certimate.me/zh-CN/docs/migrations/migrate-to-v0.4)
> - [《使用 CNAME 完成 ACME DNS-01 质询》](https://docs.certimate.me/zh-CN/blog/cname)
> - [《Why Certimate?》](https://docs.certimate.me/zh-CN/blog/why-certimate)

## 🖥️ 运行界面

[![Screenshot](https://i.imgur.com/4DAUKEE.gif)](https://www.bilibili.com/video/BV1xockeZEm2)

## 🤝 参与贡献

Certimate 是一个免费且开源的项目。我们欢迎任何人为 Certimate 做出贡献，以帮助改善 Certimate。包括但不限于：提交代码、反馈缺陷、交流想法，或分享你基于 Certimate 的使用案例。同时，我们也欢迎用户在个人博客或社交媒体上分享 Certimate。

如果你想要贡献代码，请先阅读我们的[贡献指南](./CONTRIBUTING.md)。

请在 https://github.com/certimate-go/certimate 提交 [Issues](https://github.com/certimate-go/certimate/issues) 和 [Pull Requests](https://github.com/certimate-go/certimate/pulls)。

#### 感谢以下贡献者对 Certimate 做出的贡献：

[![Contributors](https://contrib.rocks/image?repo=certimate-go/certimate)](https://github.com/certimate-go/certimate/graphs/contributors)

## ⛔ 免责声明

Certimate 遵循 [MIT License](https://opensource.org/licenses/MIT) 开源协议，完全免费提供，旨在“按现状”供用户使用。作者及贡献者不对使用本软件所产生的任何直接或间接后果承担责任，包括但不限于性能下降、数据丢失、服务中断、或任何其他类型的损害。

**无任何保证**：本软件不提供任何明示或暗示的保证，包括但不限于对特定用途的适用性、无侵权性、商用性及可靠性的保证。

**用户责任**：使用本软件即表示您理解并同意承担由此产生的一切风险及责任。

## 🌐 加入社群

- [Telegram](https://t.me/+ZXphsppxUg41YmVl)
- 微信群聊（因微信自身限制需群主邀请，可先加 [@usual2970](https://github.com/usual2970) 好友）

  <img src="https://i.imgur.com/8xwsLTA.png" width="200"/>

## ⭐ 星标趋势

在 GitHub 上为 Certimate 添加 Star 星标关注，即可第一时间获取新版本发布通知。

[![Stargazers over time](https://starchart.cc/certimate-go/certimate.svg?variant=adaptive)](https://starchart.cc/certimate-go/certimate)
