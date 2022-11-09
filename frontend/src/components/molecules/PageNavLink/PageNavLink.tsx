import { FC, ReactNode } from 'react';
import Link from 'next/link';

import * as S from './PageNavLink.styles';

export interface Props {
  dst: string;
  hasIndicator?: boolean;
  active?: boolean;
  children: ReactNode;
}

const PageNavLink: FC<Props> = ({
  dst,
  hasIndicator,
  active,
  children,
  ...props
}) => {
  return (
    <S.Wrapper {...props}>
      <Link href={dst}>{children}</Link>
      {hasIndicator && active && <span>x</span>}
    </S.Wrapper>
  );
};

export default PageNavLink;
