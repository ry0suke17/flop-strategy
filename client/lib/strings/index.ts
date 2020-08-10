export const first = (input: string | string[]): string => {
  return Array.isArray(input) ? input[0] : input;
};
