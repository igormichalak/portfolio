import { FC } from 'react';

import * as S from './SideBar.styles';

export interface Props {}

const SideBar: FC<Props> = ({ ...props }) => {
  return <S.Wrapper {...props}>Side bar</S.Wrapper>;
};

export default SideBar;
