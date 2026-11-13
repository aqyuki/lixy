import { DefaultTheme, defineConfig } from 'vitepress'
import { commentAnchorPlugin } from './plugins/comment-anchor'
import { withMermaid } from 'vitepress-plugin-mermaid'

export default withMermaid(
  defineConfig({
    srcDir: './src',
    lang: 'ja-JP',
    title: 'Lixy',
    description: 'Discord Bot Lixyの利用ガイドと開発者向け資料',

    lastUpdated: true,
    cleanUrls: true,

    vite: {
      optimizeDeps: {
        // Prebundle Mermaid's CommonJS dependencies (including fastdom) in dev.
        include: ['mermaid'],
      },
    },

    markdown: {
      config(md) {
        md.use(commentAnchorPlugin)
      },
    },

    themeConfig: {
      outline: {
        level: [2, 3],
      },
      socialLinks: socialLinks(),
      nav: nav(),
      sidebar: {
        '/guide/': { base: '/guide/', items: sidebarGuide() },
        '/development/': { base: '/development/', items: sidebarDevelopment() },
      },
      search: {
        provider: 'local',
      },
    },
  }),
)

function socialLinks(): DefaultTheme.SocialLink[] {
  return [
    {
      icon: 'github',
      link: 'https://github.com/aqyuki/lixy',
    },
  ]
}

function nav(): DefaultTheme.NavItem[] {
  return [
    {
      text: '利用ガイド',
      link: '/guide/what-is-lixy',
      activeMatch: '/guide/',
    },
    {
      text: '開発者向け',
      link: '/development/getting-started',
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
      text: 'Development',
      collapsed: false,
      items: [
        {
          text: 'Getting Started',
          link: 'getting-started',
        },
        {
          text: 'ローカルでの開発',
          link: 'local-development',
        },
        {
          text: '設定と環境変数',
          link: 'configuration',
        },
        {
          text: 'テストと品質チェック',
          link: 'testing',
        },
        {
          text: 'アーキテクチャ',
          link: 'architecture',
        },
      ],
    },
    {
      text: 'Guideline',
      collapsed: false,
      base: '/development/guideline/',
      items: [
        {
          text: 'コーディング規約',
          link: 'coding',
        },
        {
          text: 'テスト規約',
          link: 'testing',
        },
      ],
    },
    {
      text: 'Specification',
      collapsed: false,
      base: '/development/spec/',
      items: [
        {
          text: '設定関係',
          items: [{ text: '設定の読み込み', link: 'load-config' }],
        },
      ],
    },
    {
      text: 'ADR',
      collapsed: false,
      base: '/development/adr/',
      items: [
        {
          text: 'Overview',
          link: 'overview',
        },
      ],
    },
  ]
}
