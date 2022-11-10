import { FC, ReactNode } from 'react';

import * as S from './PageContainer.styles';
import { variants } from './PageContainer.motion';

export interface Props {
  children: ReactNode;
}

const PageContainer: FC<Props> = ({ children, ...props }) => {
  return (
    <S.Wrapper
      initial="hiddenLeft"
      animate="visible"
      exit="hiddenRight"
      variants={variants}
      {...props}
    >
      {children}
    </S.Wrapper>
  );
};

export default PageContainer;
