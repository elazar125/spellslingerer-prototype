export default {
  Default: {
    id: 'Default',
    name: 'Homunculus',
    image: 'Homunculus.png',
    colour: ['C'],
    ability: {},
    health: 0,
    splashColours: [],
    cardRules: [
      ({ card }) => card.type !== 'Land',
    ],
    landRules: [
      ({ card }) => card.type === 'Land',
    ],
  },
  CN: {
    id: 'CN',
    name: 'Chandra',
    image: 'Chandra.png',
    colour: ['R'],
    ability: {
      'Torch!': 'Before the game starts, deal 5 damage to your opponent.',
      'Splash of Color': 'Your deck can have up to 6 cards of a second color.',
    },
    health: 26,
    splashColours: ['C', 'W', 'U', 'B', 'G'],
    cardRules: [
      ({ card }) => card.type !== 'Land',
      ({ card, splash }) => card.colour.every((colour) => ['R', 'C', splash].includes(colour)),
      ({ card }) => ['CS', 'MC', 'CN'].includes(card.set),
    ],
    landRules: [
      ({ card }) => card.type === 'Land',
      ({ card }) => card.colour.every((colour) => ['R', 'C'].includes(colour)),
    ],
  },
  JB: {
    id: 'JB',
    name: 'Jace',
    image: 'Jace.png',
    colour: ['U'],
    ability: {
      'Multiverse Mind': 'Activate: Change the cost of each card in your deck to 1, then draw a card. This ability costs 20 mana. Reduce the cost by 1 each time you draw a card.',
      'Splash of Color': 'Your deck can have up to 6 cards of a second color.',
    },
    health: 27,
    splashColours: ['C', 'W', 'B', 'R', 'G'],
    cardRules: [
      ({ card }) => card.type !== 'Land',
      ({ card, splash }) => card.colour.every((colour) => ['U', 'C', splash].includes(colour)),
      ({ card }) => ['CS', 'MC', 'JB'].includes(card.set),
    ],
    landRules: [
      ({ card }) => card.type === 'Land',
      ({ card }) => card.colour.every((colour) => ['U', 'C'].includes(colour)),
    ],
  },
  LV: {
    id: 'LV',
    name: 'Liliana',
    image: 'Liliana.png',
    colour: ['B'],
    ability: {
      'Grave Ambition': 'After you return a creature fromm any graveyard to your hand or your side of the arena, give it +1/+1 and heal yourself for 1.',
      'Splash of Color': 'Your deck can have up to 6 cards of a second color.',
    },
    health: 26,
    splashColours: ['C', 'W', 'U', 'R', 'G'],
    cardRules: [
      ({ card }) => card.type !== 'Land',
      ({ card, splash }) => card.colour.every((colour) => ['B', 'C', splash].includes(colour)),
      ({ card }) => ['CS', 'MC', 'LV'].includes(card.set),
    ],
    landRules: [
      ({ card }) => card.type === 'Land',
      ({ card }) => card.colour.every((colour) => ['B', 'C'].includes(colour)),
    ],
  },
  GJ: {
    id: 'GJ',
    name: 'Gideon',
    image: 'Gideon.png',
    colour: ['W'],
    ability: {
      'Lead the Charge': 'Gideon can join 3 or more creatures in your attack. Drag Gideon out to summon him and launch the attack.',
      'Splash of Color': 'Your deck can have up to 6 cards of a second color.',
    },
    health: 30,
    splashColours: ['C', 'U', 'B', 'R', 'G'],
    cardRules: [
      ({ card }) => card.type !== 'Land',
      ({ card, splash }) => card.colour.every((colour) => ['W', 'C', splash].includes(colour)),
      ({ card }) => ['CS', 'MC', 'GJ'].includes(card.set),
    ],
    landRules: [
      ({ card }) => card.type === 'Land',
      ({ card }) => card.colour.every((colour) => ['W', 'C'].includes(colour)),
    ],
  },
  NR: {
    id: 'NR',
    name: 'Nissa',
    image: 'Nissa.png',
    colour: ['G'],
    ability: {
      Worldwaker: 'Your deck can have up to 6 cards of other colors and a land of any other.',
      Thrive: 'After you get a mana gem that\'s not fromm your land, heal yourself for 1.',
    },
    health: 27,
    splashColours: [],
    cardRules: [
      ({ card }) => card.type !== 'Land',
      ({ card }) => ['CS', 'MC', 'NR'].includes(card.set),
    ],
    landRules: [
      ({ card }) => card.type === 'Land',
    ],
  },
  DR: {
    id: 'DR',
    name: 'Domri',
    image: 'Domri.png',
    colour: ['R', 'G'],
    ability: {
      'Unbridled Rage': 'After you summon a creature, gain 1 rage. Spend 5 rage to summon a War Boar.',
    },
    health: 28,
    splashColours: [],
    cardRules: [
      ({ card }) => card.type !== 'Land',
      ({ card }) => card.colour.every((colour) => ['R', 'G', 'C'].includes(colour)),
      ({ card }) => ['CS', 'MC', 'DR'].includes(card.set),
    ],
    landRules: [
      ({ card }) => card.type === 'Land',
      ({ card }) => card.colour.every((colour) => ['R', 'G', 'C'].includes(colour)),
    ],
  },
  VR: {
    id: 'VR',
    name: 'Vraska',
    image: 'Vraska.png',
    colour: ['B', 'G'],
    ability: {
      'Dramatic Exit': 'During your turn, Finales of friendly creatures happen twice.',
    },
    health: 28,
    splashColours: [],
    cardRules: [
      ({ card }) => card.type !== 'Land',
      ({ card }) => card.colour.every((colour) => ['B', 'G', 'C'].includes(colour)),
      ({ card }) => ['CS', 'MC', 'VR'].includes(card.set),
    ],
    landRules: [
      ({ card }) => card.type === 'Land',
      ({ card }) => card.colour.every((colour) => ['B', 'G', 'C'].includes(colour)),
    ],
  },
  KY: {
    id: 'KY',
    name: 'Kaya',
    image: 'Kaya.png',
    colour: ['W', 'B'],
    ability: {
      'Free Spirit': 'Mission: Deal 4 damage to an opponent during your turn. If you complete this mission, add a random Unchained Spirit to your hand. Your next mission requires 1 more damage to complete.',
    },
    health: 27,
    splashColours: [],
    cardRules: [
      ({ card }) => card.type !== 'Land',
      ({ card }) => card.colour.every((colour) => ['W', 'B', 'C'].includes(colour)),
      ({ card }) => ['CS', 'MC', 'KY'].includes(card.set),
    ],
    landRules: [
      ({ card }) => card.type === 'Land',
      ({ card }) => card.colour.every((colour) => ['W', 'B', 'C'].includes(colour)),
    ],
  },
  KA: {
    id: 'KA',
    name: 'Kiora',
    image: 'Kiora.png',
    colour: ['U', 'G'],
    ability: {
      'Beckon Behemoths': 'Start the game with a Fish. After you have 10, 15, and 20 mana gems, add a random Leviathan to your hand.',
    },
    health: 28,
    splashColours: [],
    cardRules: [
      ({ card }) => card.type !== 'Land',
      ({ card }) => card.colour.every((colour) => ['U', 'G', 'C'].includes(colour)),
      ({ card }) => ['CS', 'MC', 'KA'].includes(card.set),
    ],
    landRules: [
      ({ card }) => card.type === 'Land',
      ({ card }) => card.colour.every((colour) => ['U', 'G', 'C'].includes(colour)),
    ],
  },
  AK: {
    id: 'AK',
    name: 'Ashiok',
    image: 'Ashiok.png',
    colour: ['U', 'B'],
    ability: {
      'Weave Nightmares': 'After you add a card to you hand that did not start in your deck, or play one of those cards, attach a Nightmare to a random card in your opponent\'s deck.',
    },
    health: 25,
    splashColours: [],
    cardRules: [
      ({ card }) => card.type !== 'Land',
      ({ card }) => card.colour.every((colour) => ['U', 'B', 'C'].includes(colour)),
      ({ card }) => ['CS', 'MC', 'AK'].includes(card.set),
    ],
    landRules: [
      ({ card }) => card.type === 'Land',
      ({ card }) => card.colour.every((colour) => ['U', 'B', 'C'].includes(colour)),
    ],
  },
  AG: {
    id: 'AG',
    name: 'Ajani',
    image: 'Ajani.png',
    colour: ['W', 'G'],
    ability: {
      Mentor: 'After you target a friendly creature, give the next creature you draw +1/+0.',
    },
    health: 30,
    splashColours: [],
    cardRules: [
      ({ card }) => card.type !== 'Land',
      ({ card }) => card.colour.every((colour) => ['W', 'G', 'C'].includes(colour)),
      ({ card }) => ['CS', 'MC', 'AG'].includes(card.set),
    ],
    landRules: [
      ({ card }) => card.type === 'Land',
      ({ card }) => card.colour.every((colour) => ['W', 'G', 'C'].includes(colour)),
    ],
  },
  AN: {
    id: 'AN',
    name: 'Angrath',
    image: 'Angrath.png',
    colour: ['B', 'R'],
    ability: {
      'Hammer Home': 'After you summon a creature, give it +1/+0 if your opponent was dealt damage this turn.',
    },
    health: 29,
    splashColours: [],
    cardRules: [
      ({ card }) => card.type !== 'Land',
      ({ card }) => card.colour.every((colour) => ['B', 'R', 'C'].includes(colour)),
      ({ card }) => ['CS', 'MC', 'AN'].includes(card.set),
    ],
    landRules: [
      ({ card }) => card.type === 'Land',
      ({ card }) => card.colour.every((colour) => ['B', 'R', 'C'].includes(colour)),
    ],
  },
  RZ: {
    id: 'RZ',
    name: 'Ral',
    image: 'Ral.png',
    colour: ['U', 'R'],
    ability: {
      'Storm Conduit': 'Every third spell you cast during a turn has its cost reduced by 2.',
    },
    health: 26,
    splashColours: [],
    cardRules: [
      ({ card }) => card.type !== 'Land',
      ({ card }) => card.colour.every((colour) => ['U', 'R', 'C'].includes(colour)),
      ({ card }) => ['CS', 'MC', 'RZ'].includes(card.set),
    ],
    landRules: [
      ({ card }) => card.type === 'Land',
      ({ card }) => card.colour.every((colour) => ['U', 'R', 'C'].includes(colour)),
    ],
  },
  TF: {
    id: 'TF',
    name: 'Teferi',
    image: 'Teferi.png',
    colour: ['W', 'U'],
    ability: {
      Outwit: 'After every third trap you cast, draw a card.',
    },
    health: 27,
    splashColours: [],
    cardRules: [
      ({ card }) => card.type !== 'Land',
      ({ card }) => card.colour.every((colour) => ['W', 'U', 'C'].includes(colour)),
      ({ card }) => ['CS', 'MC', 'TF'].includes(card.set),
    ],
    landRules: [
      ({ card }) => card.type === 'Land',
      ({ card }) => card.colour.every((colour) => ['W', 'U', 'C'].includes(colour)),
    ],
  },
  NH: {
    id: 'NH',
    name: 'Nahiri',
    image: 'Nahiri.png',
    colour: ['W', 'R'],
    ability: {
      'Forge From Stone': 'Spend 1 mana to create the Stoneforged Blade and give it to a friendly creature or move it to another friendly creature.',
    },
    health: 26,
    splashColours: [],
    cardRules: [
      ({ card }) => card.type !== 'Land',
      ({ card }) => card.colour.every((colour) => ['W', 'R', 'C'].includes(colour)),
      ({ card }) => ['CS', 'MC', 'NH'].includes(card.set),
    ],
    landRules: [
      ({ card }) => card.type === 'Land',
      ({ card }) => card.colour.every((colour) => ['W', 'R', 'C'].includes(colour)),
    ],
  },
};
