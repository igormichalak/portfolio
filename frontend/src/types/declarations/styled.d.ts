import 'styled-components';

import { CustomTheme } from 'utils/styled/theme';

declare module 'styled-components' {
  export interface DefaultTheme extends CustomTheme {}
}
