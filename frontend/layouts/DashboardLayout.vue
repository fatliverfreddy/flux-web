<template>
  <div class="container">
    <div class="main-wrapper">
      <nav id="slide-menu" v-click-outside="clickOutsideMenu">
        <div class="logo">
          <div>
            <img src="@/assets/images/logo.png" />
          </div>
          <div>
            <span>
              <a :href="githubUrl" target="__blank">Flux Web v{{ fluxWebRelease }}</a>
            </span>
          </div>
        </div>
        <ul>
          <a href="/">
            <li>
              <i class="fas fa-cogs"></i>Workloads
            </li>
          </a>
          <li>Events</li>
          <li>Calendar</li>
          <a href="about">
            <li class="sep">
              <i class="fas fa-question"></i>About
            </li>
          </a>
          <li>Logout</li>
        </ul>
      </nav>
      <div id="content">
        <button class="menu-trigger" @click="toggleMenu" ref="menu_btn">
          <i class="fas fa-bars"></i>
        </button>
        <div class="right-side">
          <slot name="right-side"></slot>
        </div>
      </div>
    </div>
  </div>
</template>


<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import Navbar from "~/components/Navbar.vue";
import Config from "~/config";

@Component({
  methods: {
    clickOutsideMenu(event) {
      var body = document.body;
      let menuBtn = this.$refs.menu_btn as Element;

      if (event.target != menuBtn && event.target.parentElement != menuBtn) {
        body.classList.remove("menu-active");
      }
    }
  }
})
export default class CallScreenLayout extends Vue {
  public fluxWebRelease = Config.config.fluxWebRelease;
  public githubUrl = Config.config.githubUrl;

  public toggleMenu(event: any) {
    var body = document.body;

    body.classList.toggle("menu-active");
  }
}
</script>

<style scoped lang="scss">
@import "../assets/scss/include";

/*
  Menu
*/
nav#slide-menu {
  position: fixed;
  top: 0;
  left: -100px;
  bottom: 0;
  display: block;
  float: left;
  width: 100%;
  height: 100%;
  max-width: 284px;
  background: #343d5d;

  -moz-transition: all 300ms;
  -webkit-transition: all 300ms;
  transition: all 300ms;

  .logo {
    width: 100%;
    padding-top: 15px;
    padding-left: 40px;

    img {
      width: 50%;
    }

    span {
      font-family: "Ubuntu", sans-serif;
      font-size: 13px;
      color: #fff;
      a {
        color: #fff;
        text-decoration: none;
        &:hover {
          color: #ff6a58;
        }
      }
    }
  }

  > ul {
    display: block;
    margin: 40px;
    padding: 0;
    list-style: none;
    opacity: 0.5;
    font-family: "Ubuntu", sans-serif;

    -moz-transition: all 300ms;
    -webkit-transition: all 300ms;
    transition: all 300ms;

    a:active {
      text-decoration: none;
    }

    li {
      color: #bfcddbc9;
      padding: 6px 0;
      cursor: pointer;
      transition: 0.2s;
      user-select: none;
      i {
        padding-right: 10px;
      }
      &:hover {
        color: #d0dbe6;
        transition: 0s;
        i.fas {
          color: #ffd000e0;
        }
      }
      &:active {
        color: #fff;
      }
    }

    li.sep {
      margin-top: 7px;
      padding-top: 14px;
      border-top: 1px solid lighten(#556270, 10%);
    }
  }
}

body.menu-active nav#slide-menu {
  left: 0px;
}
body.menu-active nav#slide-menu ul {
  left: 0px;
  opacity: 1;
}

/*
  Content
*/

div#content {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  overflow: auto;
  border-radius: 0;
  background: #fff;

  -moz-box-sizing: border-box;
  -webkit-box-sizing: border-box;
  box-sizing: border-box;

  -moz-box-shadow: -3px 0 6px darken(#556270, 5%);
  -webkit-box-shadow: -3px 0 6px darken(#556270, 5%);
  box-shadow: -3px 0 6px darken(#556270, 5%);

  -moz-transition: all 300ms;
  -webkit-transition: all 300ms;
  transition: all 300ms;

  button.menu-trigger {
    position: fixed;
    top: 15px;
    left: 15px;
    width: 42px;
    height: 42px;
    background: #5c6484;
    color: #fff;
    z-index: 1;
    cursor: pointer;
    border-radius: 5px;
    font-size: 16px;
    border: none;

    i {
      position: absolute;
      top: 50%;
      left: 50%;
      -moz-transform: translateX(-50%) translateY(-50%);
      -webkit-transform: translateX(-50%) translateY(-50%);
      transform: translateX(-50%) translateY(-50%);
    }

    -moz-transition: all 300ms;
    -webkit-transition: all 300ms;
    transition: all 300ms;

    &:hover {
      background: #5c6484;
      box-shadow: 0 3px 6px 0 rgba(0, 0, 0, 0.4),
        0 3px 9px 0 rgba(0, 0, 0, 0.19);
      cursor: pointer;
    }
    &:active {
      background: #5c6484;
      box-shadow: inset 0 6px 15px 0 rgba(0, 0, 0, 0.4),
        inset 0 3px 9px 0 rgba(0, 0, 0, 0.25);
      cursor: pointer;
    }
    &:focus {
      outline: 0;
    }
  }
}

body.menu-active div#content {
  left: 284px;
  border-radius: 7px 0 0 7px;
}
body.menu-active div#content .menu-trigger {
  left: 299px;

  background: #5c6484;
  box-shadow: inset 0 3px 12px 0 rgba(0, 0, 0, 0.4),
    inset 0 3px 6px 0 rgba(0, 0, 0, 0.25);
  &:hover {
    background: #727a9b;
    box-shadow: inset 0 3px 12px 0 rgba(0, 0, 0, 0.4),
      inset 0 3px 6px 0 rgba(0, 0, 0, 0.25);
  }
  &:active {
    background: #5c6484;
    box-shadow: inset 0 6px 15px 0 rgba(0, 0, 0, 0.4),
      inset 0 3px 6px 0 rgba(0, 0, 0, 0.25);
    cursor: pointer;
  }
}

.container {
  height: 100%;
  position: fixed;
  width: 100%;
  height: 100%;
  left: 0;
  top: 0;
  @include respond-above("Lmedium") {
    flex-direction: column;
  }

  .main-wrapper {
    display: flex;
    flex-direction: row;
    background: #343d5d;
    justify-content: center;
    z-index: 9999;
    position: relative;
    height: 100%;
    min-height: 450px;
    @include respond-above("Lmedium") {
      display: block;
    }
    .left-side {
      @include respond-above("Lmedium") {
        height: 60px;
        width: 100%;
      }
      width: 60px;
      background: rgb(38, 50, 80);
      background: linear-gradient(
        180deg,
        rgba(38, 50, 80, 1) 0%,
        rgba(59, 68, 106, 1) 100%
      );
    }
    .right-side {
      position: relative;
      background: #fff;
      width: 100%;
      height: 100%;
      padding: 15px;
      box-sizing: border-box;
    }
  }
}
</style>
