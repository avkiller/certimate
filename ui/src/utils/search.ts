export const matchSearchString = (keyword: string, candidate: string) => {
  keyword = String(keyword ?? "").toLowerCase();
  candidate = String(candidate ?? "").toLowerCase();

  if (keyword.length === 0) {
    return false;
  }

  if (candidate.includes(keyword)) {
    return true;
  }

  if (keyword.includes(" ")) {
    keyword = keyword.replaceAll(" ", "");
    candidate = candidate.replaceAll(" ", "");
    if (matchSearchString(keyword, candidate)) {
      return true;
    }
  }

  return false;
};

export const matchSearchOption = (keyword: string, candidate: string | { label?: unknown } | { value?: unknown }) => {
  if (typeof candidate === "string") {
    return matchSearchString(keyword, candidate);
  }

  if ("label" in candidate && candidate.label != null) {
    if (matchSearchString(keyword, candidate.label as string)) {
      return true;
    }
  }

  if ("value" in candidate && candidate.value != null) {
    if (matchSearchString(keyword, candidate.value as string)) {
      return true;
    }
  }

  return false;
};
