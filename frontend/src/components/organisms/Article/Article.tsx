import { FC } from 'react';

import { BlogPost } from 'types/BlogPost';
import Heading from 'components/atoms/Heading/Heading';

import * as S from './Article.styles';

export interface Props {
  article: BlogPost;
}

const Article: FC<Props> = ({ article, ...props }) => {
  return (
    <S.Wrapper {...props}>
      <Heading>{article.title}</Heading>
    </S.Wrapper>
  );
};

export default Article;
