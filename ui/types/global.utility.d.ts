declare global {
  type Nullish<T> = {
    [P in keyof T]?: T[P] | null | undefined;
  };

  type ArrayElement<T> = T extends (infer U)[] ? U : never;
}

export {};
