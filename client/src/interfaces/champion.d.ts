export interface Field {
  name: string;
  image_url: string;
}

export interface Rune {
  primary: {
    type: Field;
    main: Field;
    more: Field[];
  };
  secondary: {
    type: Field;
    more: Field[];
  };
  boosts: ({
    description: string;
  } & Field)[];
}

export interface Champion {
  id: string;
  name: string;
  image_url: string;
  line: string;
  skills: Field[];
  skill_order: string[];
  summoners: Field[];
  items: Field[][];
  runes: Rune[];
}
