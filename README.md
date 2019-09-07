# Comics maker

Overlay text on a template image

Example
```bash
    ./comics-maker --image template.png make
```

Config
```yaml
config:
  debug: true #draw a rectangle outline
  size: 15 # common font size
  spacing: 2.5 # common font spacing
  blocks:
    - x1: 80 #rectangle coordinates
      y1: 55
      x2: 470
      y2: 90
      size: 40 #font size
      text: "開発が始まります.." # words that do not fit in the rectangle are wrapped in a new line
    - x1: 50
      y1: 255
      x2: 200
      y2: 355
      size: 14
      text: "難しいです どうすれば いいですか？ とても難しい うまくいかないと思う とにかくやってみます"
```