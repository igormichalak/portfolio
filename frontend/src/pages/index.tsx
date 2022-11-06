import type { FC } from 'react';
import type { InferGetServerSidePropsType } from 'next';
import Head from 'next/head';
import Link from 'next/link';

import type { BlogPostFeedEntry } from 'types/BlogPostFeedEntry';
import { api } from 'utils/wretch';

export const getServerSideProps = async () => {
  const { posts } = (await api.get('/blog/feed')) as {
    posts: BlogPostFeedEntry[];
  };

  return {
    props: {
      postEntries: posts,
    },
  };
};

export interface HomePageProps
  extends InferGetServerSidePropsType<typeof getServerSideProps> {}

const Home: FC<HomePageProps> = ({ postEntries }) => {
  return (
    <div>
      <Head>
        <title>Create Next App</title>
        <meta name="description" content="Generated by create next app" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main>
        <h1>Next.js app</h1>
        <ul>
          {postEntries.map(postEntry => (
            <li key={postEntry.id}>
              <Link href={'/blog/' + postEntry.slug}>{postEntry.title}</Link>
              <p>{new Date(postEntry.updated_at).toLocaleDateString()}</p>
            </li>
          ))}
        </ul>
      </main>
    </div>
  );
};

export default Home;
