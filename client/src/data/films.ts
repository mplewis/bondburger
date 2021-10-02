import arrayShuffle from "array-shuffle";

const filmsDataPath = "films.gz";

export type Film = {
  title: string;
  year: string;
  actor: string;
  plot: string[];
};
export type Slug = string;
export type FilmsBySlug = { [slug: Slug]: Film };

var _films: FilmsBySlug = {};

export async function films(): Promise<FilmsBySlug> {
  if (Object.keys(_films).length === 0) {
    const resp = await fetch(filmsDataPath);
    const json = await resp.text();
    _films = JSON.parse(json);
  }
  return _films;
}

export async function pick(n: number): Promise<FilmsBySlug> {
  const f = await films();
  const shufKeys = arrayShuffle(Object.keys(f));
  const keys = shufKeys.slice(0, n);
  const result: FilmsBySlug = {};
  for (const key of keys) {
    result[key] = f[key];
  }
  return result;
}
