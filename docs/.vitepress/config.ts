import { DefaultTheme, defineConfig } from 'vitepress'
import { commentAnchorPlugin } from './plugins/comment-anchor'

export default defineConfig({
  srcDir: './src',
  lang: 'ja-JP',
  title: 'Lixy',
  description: 'Discord Bot Lixyの利用ガイドと開発者向け資料',
  markdown: {
    config(md) {
      md.use(commentAnchorPlugin)
    },
  },
  themeConfig: {
    nav: nav(),
    sidebar: {
      '/guide/': { base: '/guide/', items: sidebarGuide() },
      '/development/': { base: '/development/', items: sidebarDevelopment() },
    },
  },
})

function nav(): DefaultTheme.NavItem[] {
  return [
    {
      text: '利用ガイド',
      link: '/guide/what-is-lixy',
      activeMatch: '/guide/',
    },
    {
      text: '開発者向け',
      link: '/development/',
      activeMatch: '/development/',
    },
  ]
}

function sidebarGuide(): DefaultTheme.SidebarItem[] {
  return [
    {
      text: 'Introduction',
      collapsed: false,
      items: [
        {
          text: 'What is Lixy?',
          link: 'what-is-lixy',
        },
      ],
    },
  ]
}

function sidebarDevelopment(): DefaultTheme.SidebarItem[] {
  return [
    {
      text: 'Introduction',
      collapsed: false,
      items: [
        {
          text: 'はじめに',
          link: 'index.md',
        },
      ],
    },
    {
      text: 'ADR',
      collapsed: false,
      base: '/development/adr/',
      items: [
        {
          text: 'ADRについて',
          link: 'about-adr',
        },
      ],
    },
  ]
}
