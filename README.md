# Comics maker

### Script overlay text blocks on the template.

We superimpose text blocks on the template image downloaded from the link or downloaded from the local file (depending on the parameter passed).

If the text does not fit in width, then it is wrapped to the next line. If there is not enough vertical space, then the text outside the block is no longer displayed.

In the configuration, you can specify general values that will be applied to all text blocks if they are not overridden in them.

Debug mode displays red rectangles around each text block.

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
  textAlign: "left" # common text align
  blocks:
    - x1: 80 #rectangle coordinates
      y1: 55
      x2: 470
      y2: 90
      size: 40 #font size
      textAlign: "center"
      text: "開発が始まります.." # words that do not fit in the rectangle are wrapped in a new line
    - x1: 50
      y1: 255
      x2: 200
      y2: 355
      size: 14
      text: "難しいです どうすれば いいですか？ とても難しい うまくいかないと思う とにかくやってみます"
```