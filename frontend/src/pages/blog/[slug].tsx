import type { FC } from 'react';
import type { GetServerSideProps, InferGetServerSidePropsType } from 'next';
import Head from 'next/head';

import { BlogPost } from 'types/BlogPost';
import { api } from 'utils/wretch';
import { slugRegExp } from 'utils/regExp';
import PageContainer from 'components/atoms/PageContainer/PageContainer';

type Data = {
  post: BlogPost;
};

export const getServerSideProps: GetServerSideProps<Data> = async ctx => {
  const slug = (ctx.params?.slug as string) || '';
  const isValidSlug = slugRegExp.test(slug);

  if (!isValidSlug) {
    return {
      notFound: true,
    };
  }

  let post;
  try {
    const data = (await api.get(`/blog/post/${slug}`).json()) as Data;
    post = data?.post;
  } catch {
    return {
      notFound: true,
    };
  }

  return {
    props: {
      post,
    },
  };
};

export interface BlogPostPageProps
  extends InferGetServerSidePropsType<typeof getServerSideProps> {}

const BlogPost: FC<BlogPostPageProps> = ({ post }) => {
  const title = `${
    post?.title ? `${post?.title} - ` : ``
  }The Igor Michalak Blog`;

  return (
    <PageContainer>
      <Head>
        <title>{title}</title>
        <meta name="description" content="Generated by create next app" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <h1>{post?.title}</h1>
      <span dangerouslySetInnerHTML={{ __html: post?.body }} />
    </PageContainer>
  );
};

export default BlogPost;
