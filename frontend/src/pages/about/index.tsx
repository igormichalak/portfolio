import type { FC } from 'react';
import Head from 'next/head';

import PageContainer from 'components/atoms/PageContainer/PageContainer';

export interface AboutPageProps {}

const About: FC<AboutPageProps> = () => {
  return (
    <PageContainer>
      <Head>
        <title>About Igor Michalak</title>
        <meta name="description" content="Generated by create next app" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <h1>About</h1>
    </PageContainer>
  );
};

export default About;