export type FormState = {
  error?: string;
  fields?: Record<string, string>;
  blockers?: string[];
  ok?: string;
  /** What to put back in the form, which React resets after each action. */
  values?: Record<string, string>;
};
