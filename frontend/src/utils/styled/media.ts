const breakpoints = {
  xs: 320,
  s: 414,
  m: 768,
  l: 1024,
  xl: 1280,
  xxl: 1440,
  max: 1680,
};

const customMediaQuery = (minWidth: number) =>
  `@media screen and (min-width: ${minWidth}px)`;

const media = {
  xs: customMediaQuery(breakpoints.xs),
  s: customMediaQuery(breakpoints.s),
  m: customMediaQuery(breakpoints.m),
  l: customMediaQuery(breakpoints.l),
  xl: customMediaQuery(breakpoints.xl),
  xxl: customMediaQuery(breakpoints.xxl),
  max: customMediaQuery(breakpoints.max),
  reducedMotion: `@media (prefers-reduced-motion)`,
};

export { breakpoints };
export default media;
