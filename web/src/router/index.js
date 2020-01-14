import Vue from "vue";
import VueRouter from "vue-router";

Vue.use(VueRouter);

function loadView(view) {
  return () =>
    import(/* webpackChunkName: "view-[request]" */ "@/views/" + view + ".vue");
}

const routes = [
  {
    path: "/",
    name: "overview",
    component: loadView("PageOverview")
  },
  {
    path: "/objectives",
    name: "objectives",
    component: loadView("PageObjectives")
  },
  {
    path: "/key-results",
    name: "key-results",
    component: loadView("PageKeyResults")
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export default router;
