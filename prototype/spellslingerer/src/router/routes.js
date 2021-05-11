export default [
  {
    path: '/',
    inMenu: true,
    linkText: 'Home',
    icon: 'mdi-home',
    component: () => import('../views/home.vue'),
  },
  {
    path: '/cards',
    inMenu: true,
    linkText: 'Card Search',
    icon: '$vuetify.icons.cardMagnify',
    component: () => import('../views/wrapper.vue'),
    children: [
      {
        path: ':id',
        component: () => import(/* webpackChunkName: "cards" */ '../views/card.vue'),
      },
      {
        path: '',
        component: () => import(/* webpackChunkName: "cards" */ '../views/search.vue'),
      },
    ],
  },
  {
    path: '/decks',
    inMenu: true,
    nested: true,
    linkText: 'Deck Builder',
    icon: 'mdi-cards',
    component: () => import('../views/wrapper.vue'),
    children: [
      {
        path: 'search',
        inMenu: true,
        linkText: 'Deck Search',
        icon: '$vuetify.icons.cardsMagnify',
        component: () => import(/* webpackChunkName: "decks" */ '../views/decks.vue'),
      },
      {
        path: 'trending',
        linkText: 'Trending Decks',
        icon: '$vuetify.icons.cardsTrending',
      },
      {
        path: 'builder',
        inMenu: true,
        linkText: 'My Decks',
        icon: '$vuetify.icons.cardsHeart',
        component: () => import(/* webpackChunkName: "decks" */ '../views/my-decks.vue'),
      },
      {
        path: 'starred',
        linkText: 'Starred Decks',
        icon: '$vuetify.icons.cardsStar',
      },
      {
        path: ':deckId',
        inMenu: true,
        linkText: 'New Deck',
        icon: '$vuetify.icons.cardsPlus',
        component: () => import(/* webpackChunkName: "decks" */ '../views/deck-builder.vue'),
      },
    ],
  },
  {
    path: '/my',
    inMenu: true,
    nested: true,
    linkText: 'My Lists',
    icon: 'mdi-account-details',
    component: () => import('../views/wrapper.vue'),
    children: [
      {
        path: 'wishlist',
        inMenu: true,
        linkText: 'My Wishlist',
        icon: 'mdi-playlist-star',
        component: () => import(/* webpackChunkName: "lists" */ '../views/wishlist.vue'),
      },
      {
        path: 'collection',
        inMenu: true,
        linkText: 'My Collection',
        icon: 'mdi-playlist-check',
        component: () => import(/* webpackChunkName: "lists" */ '../views/collection.vue'),
      },
      {
        path: 'materials',
        inMenu: true,
        linkText: 'My Crafting Materials',
        icon: 'mdi-hammer',
        component: () => import(/* webpackChunkName: "crafting" */ '../views/crafting.vue'),
      },
    ],
  },
  {
    path: '/tier-list',
    linkText: 'Tier List',
    icon: 'mdi-chart-bar',
  },
  {
    path: '/user/:id',
    inCorner: true,
    linkText: 'My Account',
    icon: 'mdi-account-circle',
    component: () => import(/* webpackChunkName: "user" */ '../views/user.vue'),
  },
  {
    path: '/about',
    inCorner: true,
    linkText: 'About This Site',
    icon: 'mdi-help-circle',
    component: () => import(/* webpackChunkName: "about" */ '../views/about.vue'),
  },
  {
    path: '/my-data',
    inCorner: true,
    linkText: 'Manage My Data',
    icon: 'mdi-settings',
  },
  {
    path: '*',
    component: () => import('../views/error.vue'),
  },
];