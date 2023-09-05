export const pageTitle = (prefix: string): string => {
  const suffix = "SZPP Judge";
  return prefix ? `${prefix} | ${suffix}` : suffix;
};
