<template>
  <v-app>
    <v-app-bar app color="secondary" clipped-left>
      <v-app-bar-nav-icon @click.stop="drawer = !drawer" />
      <v-toolbar-title>
        <router-link to="/">
          <v-img
            src="/Spellslingerer.png"
            contain
            width="200"
            class="mb-2"
          />
        </router-link>
      </v-toolbar-title>
      <v-spacer />

      <v-tooltip bottom open-on-click :open-on-hover="false">
        <template v-slot:activator="{ on }">
          <v-btn icon @click="copyUrl()" v-on="on">
            <v-icon>mdi-share-variant</v-icon>
          </v-btn>
        </template>
        <span>URL copied to clipboard</span>
      </v-tooltip>

      <v-menu bottom offset-y>
        <template #activator="{ on }">
          <v-btn icon v-on="on">
            <v-icon>mdi-account-circle</v-icon>
          </v-btn>
        </template>

        <v-list>
          <v-list-item
            v-for="item in cornerLinks"
            :key="item.linkText"
            :to="`${item.path.startsWith(':') ? 'new' : item.path}` /* Use current user in place of ':id' placeholder on User link */"
          >
            <v-list-item-icon>
              <v-icon v-text="item.icon" />
            </v-list-item-icon>
            <v-list-item-title v-text="item.linkText" />
          </v-list-item>
        </v-list>
      </v-menu>
    </v-app-bar>

    <v-navigation-drawer app v-model="drawer" clipped :mini="!drawerPinned" :expand-on-hover="!drawerPinned">
      <v-list>
        <v-list-item v-for="item in links" :key="item.linkText" :to="item.path" color="primary">
          <v-list-item-icon>
            <v-icon v-text="item.icon" />
          </v-list-item-icon>
          <v-list-item-title v-text="item.linkText" />
        </v-list-item>

        <v-list-group
          :prepend-icon="item.icon"
          v-for="item in nestedLinks"
          :key="item.linkText"
          color="white"
        >
          <template #activator>
            <v-list-item-title v-text="item.linkText" />
          </template>

          <v-list-item
            v-for="subitem in item.children.filter(({ inMenu }) => inMenu)"
            :key="subitem.linkText"
            :to="`${item.path}/${subitem.path.startsWith(':') ? 'new' : subitem.path}` /* Use 'new' in place of ':deckId' placeholder on New Deck link */"
            color="primary"
          >
            <v-list-item-icon>
              <v-icon v-text="subitem.icon" />
            </v-list-item-icon>
            <v-list-item-title v-text="subitem.linkText" />
          </v-list-item>
        </v-list-group>
      </v-list>
      <template #prepend>
        <v-list-item dense @click="drawerPinned = !drawerPinned">
          <span><!-- Render a blank spot on mini menu --></span>
          <v-list-item-title v-text="'Keep Open'" />
          <v-icon small>mdi-pin{{ drawerPinned ? '' : '-outline' }}</v-icon>
        </v-list-item>
      </template>
      <template #append>
        <v-list-item align="center">
          <span><!-- Render a blank spot on mini menu --></span>
          <v-btn-toggle group>
            <v-btn icon href=""><v-icon>$vuetify.icons.github</v-icon></v-btn>
            <v-btn icon href=" https://discord.gg/zYD3sW8"><v-icon>mdi-discord</v-icon></v-btn>
            <v-btn icon href=""><v-icon>mdi-patreon</v-icon></v-btn>
            <v-btn icon href="mailto:spellslingererdeckbuilder@gmail.com"><v-icon>mdi-email</v-icon></v-btn>
          </v-btn-toggle>
        </v-list-item>
      </template>
    </v-navigation-drawer>

    <v-content>
      <v-container>
        <v-row>
          <v-col>
            <router-view />
          </v-col>
        </v-row>
      </v-container>
    </v-content>
  </v-app>
</template>

<script>
import routes from './router/routes';

export default {
  name: 'App',

  components: {
  },

  data: () => ({
    drawer: true,
    drawerPinned: false,
    routes,
  }),

  computed: {
    links() {
      return this.routes.filter((route) => route.linkText && !route.nested && route.inMenu);
    },
    nestedLinks() {
      return this.routes.filter((route) => route.linkText && route.nested && route.inMenu);
    },
    cornerLinks() {
      return this.routes.filter((route) => route.linkText && route.inCorner);
    },
    isMobile() {
      return this.$vuetify.breakpoint.smAndDown;
    },
  },

  methods: {
    copyLink() {
      const el = document.createElement('textarea');
      el.value = window.location.href;
      el.setAttribute('readonly', '');
      el.style.position = 'absolute';
      el.style.left = '-9999px';
      document.body.appendChild(el);
      const selected = [];
      for (let i = 0; i < document.getSelection().rangeCount; i++) {
        selected.push(document.getSelection().getRangeAt(i));
      }
      document.getSelection().removeAllRanges();
      el.select();
      document.execCommand('copy');
      document.body.removeChild(el);
      document.getSelection().removeAllRanges();
      selected.forEach((selection) => document.getSelection().addRange(selection));
    },
  },
};
</script>
