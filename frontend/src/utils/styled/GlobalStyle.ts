import { createGlobalStyle } from 'styled-components';

import media from 'utils/styled/media';

const GlobalStyle = createGlobalStyle`
  ${media.reducedMotion} {  
    html {
      scroll-behavior: smooth;
    }
  }
  
  html {
    -webkit-text-size-adjust: 100%;
    text-size-adjust: 100%;
  }
  
  body {
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-rendering: optimizeLegibility;
    font-kerning: normal;
  }
  
  *,
  *::before,
  *::after {
    box-sizing: border-box;
  }
  
  * {
    margin: 0;
    padding: 0;
  }
  
  a {
    text-decoration: none;
  }
  
  button {
    background-color: transparent;
  }
  
  a, button {
    color: inherit;
    touch-action: manipulation;
  }
  
  img {
    max-width: 100%;
  }
  
  svg {
    width: 100%;
    height: 100%;
    pointer-events: none;
  }
`;

export default GlobalStyle;
