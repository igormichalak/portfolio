import { FC } from 'react';

import NavList from 'components/organisms/NavList/NavList';

import * as S from './SideBar.styles';

export interface Props {}

const SideBar: FC<Props> = ({ ...props }) => {
  return (
    <S.Wrapper {...props}>
      <NavList />
    </S.Wrapper>
  );
};

export default SideBar;
