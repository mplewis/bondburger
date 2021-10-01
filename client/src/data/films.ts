const filmsDataPath = "films.gz";

export type Film = {
  title: string;
  year: string;
  actor: string;
  plot: string[];
};
export type Slug = string;
export type FilmsBySlug = { [slug: Slug]: Film };

var _films: FilmsBySlug | null = null;

export default async function films(): Promise<FilmsBySlug> {
  if (_films) return _films;

  const resp = await fetch(filmsDataPath);
  const json = await resp.text();
  _films = JSON.parse(json);

  if (!_films) throw new Error("Could not load films"); // type assertion
  return _films;
}
