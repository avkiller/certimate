﻿import { Fragment, cloneElement, createElement, isValidElement, useMemo } from "react";

export type UseTriggerElementOptions = {
  onClick?: (e: React.MouseEvent) => void;
};

/**
 * 获取一个触发器元素。
 * 通常为配合 Drawer、Modal 等组件使用。
 * @param {React.ReactNode} trigger
 * @param {UseTriggerElementOptions} [options]
 * @returns {React.ReactElement}
 */
const useTriggerElement = (trigger: React.ReactNode, options?: UseTriggerElementOptions) => {
  const onClick = options?.onClick;
  const triggerDom = useMemo(() => {
    if (!trigger) {
      return null;
    }

    const el = isValidElement(trigger) ? trigger : createElement(Fragment, null, trigger);
    return cloneElement(el, {
      ...el.props,
      onClick: (e: React.MouseEvent) => {
        onClick?.(e);
        el.props?.onClick?.(e);
      },
    });
  }, [trigger, onClick]);
  return triggerDom;
};

export default useTriggerElement;
