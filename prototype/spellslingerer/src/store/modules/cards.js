import allCards from '../data/cards.json';

export default {
  namespaced: true,
  getters: {
    cards: () => Object.values(allCards).filter(({ id }) => id !== 'Default'),
    card: () => (cardId) => allCards[cardId],
    cardNames: () => Object.values(allCards).filter(({ id }) => id !== 'Default').map(({ name }) => name),
    landNames: () => Object.values(allCards).filter(({ type }) => type === 'Land').map(({ name }) => name),
  },
};
