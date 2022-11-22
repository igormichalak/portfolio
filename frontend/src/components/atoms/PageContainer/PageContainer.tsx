import { FC, ReactNode } from 'react';

import * as S from './PageContainer.styles';
import { variants } from './PageContainer.motion';

export interface Props {
  topMargin?: boolean;
  centerContent?: boolean;
  children: ReactNode;
}

const PageContainer: FC<Props> = ({
  topMargin,
  centerContent,
  children,
  ...props
}) => {
  return (
    <S.Wrapper
      initial="hiddenLeft"
      animate="visible"
      exit="hiddenRight"
      variants={variants}
      $topMargin={topMargin}
      $centerContent={centerContent}
      {...props}
    >
      {children}
    </S.Wrapper>
  );
};

export default PageContainer;
