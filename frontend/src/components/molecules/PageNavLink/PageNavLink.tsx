import { FC, ReactNode } from 'react';
import Link from 'next/link';

import * as S from './PageNavLink.styles';

export interface Props {
  dst: string;
  hasIndicator?: boolean;
  isActive?: boolean;
  children: ReactNode;
}

const PageNavLink: FC<Props> = ({
  dst,
  hasIndicator,
  isActive,
  children,
  ...props
}) => {
  return (
    <S.Wrapper {...props}>
      {hasIndicator && isActive && (
        <S.ActiveIndicator layoutId="activeRouteIndicator" />
      )}
      <Link href={dst} passHref legacyBehavior>
        <S.Link isActive={!!isActive}>{children}</S.Link>
      </Link>
    </S.Wrapper>
  );
};

export default PageNavLink;
