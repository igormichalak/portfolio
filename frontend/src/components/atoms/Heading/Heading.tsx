import { FC, ReactNode } from 'react';

import * as S from './Heading.styles';

export interface Props {
  children: ReactNode;
}

const Heading: FC<Props> = ({ children, ...props }) => {
  return <S.Wrapper {...props}>{children}</S.Wrapper>;
};

export default Heading;
