export const getRange = (start: number, end: number): number[] => {
  return [...Array(end - start).keys()].map((el): number => el + start);
};
