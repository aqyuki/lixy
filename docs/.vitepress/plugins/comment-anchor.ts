import type { MarkdownRenderer } from 'vitepress'

// Run before VitePress generates permalinks and collects headings for the TOC.
export function commentAnchorPlugin(md: MarkdownRenderer) {
  md.core.ruler.before('anchor', 'comment-anchor', (state) => {
    const lines = state.src.split('\n')
    for (let index = 1; index < state.tokens.length; index++) {
      const heading = state.tokens[index]
      const comment = state.tokens[index - 1]
      if (heading.type !== 'heading_open' || comment.type !== 'html_block') continue

      const match = /^<!--\s*anchor:\s*([^\s<>]+)\s*-->\s*$/.exec(comment.content)
      if (!match) continue

      // Allow blank lines, but never skip other Markdown content.
      if (!comment.map || !heading.map) continue
      if (lines.slice(comment.map[1], heading.map[0]).some((line) => line.trim() !== '')) {
        continue
      }
      if (heading.attrGet('id') !== null) {
        throw new Error('Use either an anchor comment or a heading id, not both.')
      }

      heading.attrSet('id', match[1])
      comment.content = ''
    }
  })
}
