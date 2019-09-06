# Comics maker

Overlay text on a template image

Example
```bash
    ./comics-maker --image template.png make
```

Config
```yaml
config:
  size: 15 # common font size
  spacing: 1.5  # common font line spacing
  blocks:
    - x: 100  # horizontal coordinate of the text block
      y: 45 # horizontal coordinate of the text block
      size: 40 # block font size
      strings:
        - "開発が始まります"  # array of strings for a text block
    - x: 50
      y: 260
      size: 15
      spacing: 1.5
      strings:
        - "どうすればいいですか？"
        - "私は思うだろう.."
        - "これを試してみます"
    - x: 440
      y: 290
      strings:
        - "何も起こりません（（"
```