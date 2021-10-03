export type Film = {
  title: string;
  year: string;
  actor: string;
  plot: string[];
};
export type Slug = string;
export type FilmsBySlug = { [slug: Slug]: Film };
export type Response = { repsonse: Slug; answer: Slug };
