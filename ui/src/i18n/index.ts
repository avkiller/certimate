import i18n from "i18next";
import { initReactI18next } from "react-i18next";
import LanguageDetector from "i18next-browser-languagedetector";

import resources, { LOCALE_ZH_NAME, LOCALE_EN_NAME } from "./locales";

i18n
  .use(LanguageDetector)
  .use(initReactI18next)
  .init({
    resources,
    fallbackLng: LOCALE_ZH_NAME,
    debug: true,
    interpolation: {
      escapeValue: false,
    },
    backend: {
      loadPath: "/locales/{{lng}}.json",
    },
  });

export const localeNames = {
  ZH: LOCALE_ZH_NAME,
  EN: LOCALE_EN_NAME,
};

export default i18n;
