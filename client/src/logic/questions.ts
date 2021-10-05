import { Film } from "../types";
import { random, randInt } from "../random";
import { contextPadding } from "../config";

export type Question = {
  text: string;
  options: Film[];
  answer: Film;
};

export function buildQuestion(films: Film[]): Question {
  const answer = random(films);
  const min = contextPadding;
  const max = answer.plot.length - 1 - contextPadding;
  const base = randInt(min, max);
  const q = answer.plot.slice(base - contextPadding, base + contextPadding + 1);
  const text = q.join(" ");
  return { text, options: films, answer };
}
