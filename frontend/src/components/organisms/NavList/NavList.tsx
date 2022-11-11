import { FC } from 'react';
import { useRouter } from 'next/router';

import routes from 'utils/routes';
import PageNavLink from 'components/molecules/PageNavLink/PageNavLink';

import * as S from './NavList.styles';

export interface Props {}

const NavList: FC<Props> = ({ ...props }) => {
  const router = useRouter();

  return (
    <S.Wrapper {...props}>
      <S.LinkList>
        <PageNavLink
          dst={routes.LANDING}
          hasIndicator
          isActive={router.pathname === routes.LANDING}
        >
          Landing
        </PageNavLink>
        <PageNavLink
          dst={routes.BLOG}
          hasIndicator
          isActive={router.pathname.startsWith(routes.BLOG)}
        >
          Blog
        </PageNavLink>
        <PageNavLink
          dst={routes.ABOUT}
          hasIndicator
          isActive={router.pathname === routes.ABOUT}
        >
          About
        </PageNavLink>
      </S.LinkList>
    </S.Wrapper>
  );
};

export default NavList;
