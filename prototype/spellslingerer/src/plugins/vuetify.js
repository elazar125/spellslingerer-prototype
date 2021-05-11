import Vue from 'vue';
import Vuetify from 'vuetify/lib';
import zero from '../components/mana/0.vue';
import one from '../components/mana/1.vue';
import two from '../components/mana/2.vue';
import three from '../components/mana/3.vue';
import four from '../components/mana/4.vue';
import five from '../components/mana/5.vue';
import six from '../components/mana/6.vue';
import seven from '../components/mana/7.vue';
import eight from '../components/mana/8.vue';
import nine from '../components/mana/9.vue';
import ten from '../components/mana/10.vue';
import eleven from '../components/mana/11.vue';
import twelve from '../components/mana/12.vue';
import thirteen from '../components/mana/13.vue';
import fourteen from '../components/mana/14.vue';
import fifteen from '../components/mana/15.vue';
import white from '../components/mana/W.vue';
import blue from '../components/mana/U.vue';
import black from '../components/mana/B.vue';
import red from '../components/mana/R.vue';
import green from '../components/mana/G.vue';
import colourless from '../components/mana/C.vue';
import cardsHeart from '../components/icons/cards-heart.vue';
import cardsMagnify from '../components/icons/cards-magnify.vue';
import cardsStar from '../components/icons/cards-star.vue';
import cardsTrending from '../components/icons/cards-trending.vue';
import cardsPlus from '../components/icons/cards-plus.vue';
import cardMagnify from '../components/icons/card-magnify.vue';
import github from '../components/icons/github.vue';

Vue.component('zero', zero);
Vue.component('one', one);
Vue.component('two', two);
Vue.component('three', three);
Vue.component('four', four);
Vue.component('five', five);
Vue.component('six', six);
Vue.component('seven', seven);
Vue.component('eight', eight);
Vue.component('nine', nine);
Vue.component('ten', ten);
Vue.component('eleven', eleven);
Vue.component('twelve', twelve);
Vue.component('thirteen', thirteen);
Vue.component('fourteen', fourteen);
Vue.component('fifteen', fifteen);
Vue.component('white', white);
Vue.component('blue', blue);
Vue.component('black', black);
Vue.component('red', red);
Vue.component('green', green);
Vue.component('colourless', colourless);
Vue.component('cards-heart', cardsHeart);
Vue.component('cards-magnify', cardsMagnify);
Vue.component('cards-star', cardsStar);
Vue.component('cards-trending', cardsTrending);
Vue.component('cards-plus', cardsPlus);
Vue.component('cards-magnify', cardMagnify);
Vue.component('github', github);

Vue.use(Vuetify);

export default new Vuetify({
  theme: {
    dark: true,
    themes: {
      light: {
        primary: '#ff8200',
        secondary: '#995058',
        accent: '#8a5d32',
        error: '#f44336',
        info: '#37817b',
        success: '#5fa649',
        warning: '#dc9611',
      },
      dark: {
        primary: '#ff8200',
        secondary: '#995058',
        accent: '#8a5d32',
        error: '#f44336',
        info: '#37817b',
        success: '#5fa649',
        warning: '#dc9611',
      },
    },
  },
  icons: {
    values: {
      0: { component: zero },
      1: { component: one },
      2: { component: two },
      3: { component: three },
      4: { component: four },
      5: { component: five },
      6: { component: six },
      7: { component: seven },
      8: { component: eight },
      9: { component: nine },
      10: { component: ten },
      11: { component: eleven },
      12: { component: twelve },
      13: { component: thirteen },
      14: { component: fourteen },
      15: { component: fifteen },
      W: { component: white },
      U: { component: blue },
      B: { component: black },
      R: { component: red },
      G: { component: green },
      C: { component: colourless },
      cardsHeart: { component: cardsHeart },
      cardsMagnify: { component: cardsMagnify },
      cardsStar: { component: cardsStar },
      cardsTrending: { component: cardsTrending },
      cardsPlus: { component: cardsPlus },
      cardMagnify: { component: cardMagnify },
      github: { component: github },
    },
  },
});
