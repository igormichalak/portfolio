import type { FC } from 'react';
import type { GetServerSideProps, InferGetServerSidePropsType } from 'next';
import Head from 'next/head';
import Link from 'next/link';

import type { BlogPostFeedEntry } from 'types/BlogPostFeedEntry';
import { api } from 'utils/wretch';
import PageContainer from 'components/atoms/PageContainer/PageContainer';

type Data = {
  postEntries: BlogPostFeedEntry[];
};

export const getServerSideProps: GetServerSideProps<Data> = async () => {
  const { posts } = (await api.get('/blog/feed').json()) as {
    posts: BlogPostFeedEntry[];
  };

  return {
    props: {
      postEntries: posts,
    },
  };
};

export interface BlogFeedPageProps
  extends InferGetServerSidePropsType<typeof getServerSideProps> {}

const BlogFeed: FC<BlogFeedPageProps> = ({ postEntries }) => {
  return (
    <PageContainer>
      <Head>
        <title>The Igor Michalak Blog</title>
        <meta name="description" content="Generated by create next app" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <h1>Blog</h1>
      <ul>
        {postEntries.map(postEntry => (
          <li key={postEntry.id}>
            <Link
              href={{
                pathname: '/blog/[slug]',
                query: { slug: postEntry.slug },
              }}
            >
              {postEntry.title}
            </Link>
            <p>{new Date(postEntry.updated_at).toLocaleDateString()}</p>
          </li>
        ))}
      </ul>
    </PageContainer>
  );
};

export default BlogFeed;
