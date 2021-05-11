import Vue from 'vue';
import Vuex from 'vuex';
import materials from './modules/materials';
import wishlist from './modules/wishlist';
import collection from './modules/collection';
import decks from './modules/decks';
import cards from './modules/cards';
import planeswalkers from './modules/planeswalkers';

Vue.use(Vuex);

const store = new Vuex.Store({
  state: {
    currentUser: 'me',
  },
  modules: {
    materials,
    wishlist,
    collection,
    decks,
    cards,
    planeswalkers,
  },
});

export default store;
