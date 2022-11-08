const colors = {
  neutral: [
    'hsl(251, 24%, 8%)',
    'hsl(251, 16%, 18%)',
    'hsl(251, 8%, 28%)',
    'hsl(251, 4%, 38%)',
    'hsl(251, 4%, 50%)',
    'hsl(251, 4%, 64%)',
    'hsl(251, 8%, 74%)',
    'hsl(251, 16%, 84%)',
    'hsl(251, 24%, 94%)',
  ],
};

const zIndexes = {
  debug: 9999,
};

const theme = { colors, zIndexes };

export type CustomTheme = typeof theme;
export { colors };
export default theme;
