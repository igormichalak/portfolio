import { BlogTag } from 'types/BlogTag';

export type BlogPost = {
  id: number;
  slug: string;
  title: string;
  body: string;
  tags?: BlogTag[];
  created_at: string;
  updated_at: string;
};
