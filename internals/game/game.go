package game

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/walonCode/tetris-go/internals/block"
	"github.com/walonCode/tetris-go/internals/grid"
)

type Game struct {
	// public
	GameOver bool
	Score    int
	Music    rl.Music

	// private
	grid         grid.Grid
	blocks       []*block.Block
	currentBlock *block.Block
	nextBlock    *block.Block
	rotateSound  rl.Sound
	clearSound   rl.Sound
}

// constructor
func New() *Game {
	rl.InitAudioDevice()
	g := &Game{
		grid: *grid.New(),
		GameOver: false,
		Score: 0,
		Music: rl.LoadMusicStream("sounds/music.mp3"),
		rotateSound: rl.LoadSound("sounds/rotate.mp3"),
		clearSound: rl.LoadSound("sounds/clear.mp3"),	
	}
	rl.PauseMusicStream(g.Music)

	g.blocks = g.getAllBlocks()
	g.currentBlock = g.getRandomBlock()
	g.nextBlock = g.getRandomBlock()
	return g
}

// public functions
func (g *Game) Draw() {
	g.grid.Draw()
	g.currentBlock.Draw(11,11)
	switch g.nextBlock.Id {
		case 3:
			g.nextBlock.Draw(255,290)
		case 4:
			g.nextBlock.Draw(255,280)
		default:
			g.nextBlock.Draw(255,270)
	}
}

func (g *Game) HandleInput() {
	keyPressed := rl.GetKeyPressed()

	if g.GameOver && keyPressed != 0 {
		g.GameOver = false
		g.reset()
	}

	switch keyPressed{
		case rl.KeyLeft:
			g.moveBlockLeft()
		case rl.KeyRight:
			g.moveBlockRight()
		case rl.KeyDown:
			g.MoveBlockDown()
			g.updateScore(0,1)
		case rl.KeyUp:
			g.rotateBlock()	
	}
}

func (g *Game) MoveBlockDown() {
	if !g.GameOver {
		g.currentBlock.Move(1,0)
		if g.isBlockOutside() || g.blockFits() == false{
			g.currentBlock.Move(-1, 0)
			g.lockBlock()
		}
	}
}

// private functions
func (g *Game) moveBlockLeft() {
	if !g.GameOver {
		g.currentBlock.Move(0,-1)
		if g.isBlockOutside() || g.blockFits() == false{
			g.currentBlock.Move(0,1)
		}
	}
}

func (g *Game) moveBlockRight() {
	if !g.GameOver {
		g.currentBlock.Move(0,1)
		if g.isBlockOutside() || g.blockFits() == false{
			g.currentBlock.Move(0,-1)
		}
	}
}


func (g *Game) getRandomBlock() *block.Block {
	if len(g.blocks) == 0 {
		g.blocks = g.getAllBlocks()
	}

	randomIndex := rand.Intn(len(g.blocks))

	b := g.blocks[randomIndex]

	g.blocks = append(g.blocks[:randomIndex], g.blocks[randomIndex+1:]...)

	return b
}

func (g *Game) getAllBlocks() []*block.Block {
	return []*block.Block{
		block.NewIBlock(),
		block.NewJBlock(),
		block.NewLBlock(),
		block.NewOBlock(),
		block.NewSBlock(),
		block.NewTBlock(),
		block.NewZBlock(),
	}
}

func (g *Game) isBlockOutside() bool {
	titles := g.currentBlock.GetCellPostions()

	for _, item := range titles {
		if g.grid.IsCellOutside(item.Row, item.Column){
			return true
		}
	}

	return false
}

func (g *Game) rotateBlock() {
	if !g.GameOver {
		g.currentBlock.Rotate()
		if !g.isBlockOutside() || g.blockFits() == false {
			g.currentBlock.UndoRotation()
		}else {
			rl.PauseSound(g.rotateSound)
		}
	}
}

func (g *Game) lockBlock() {
	titles := g.currentBlock.GetCellPostions()

	for _, item := range titles {
		g.grid.Grid[item.Row][item.Column] = g.currentBlock.Id
	}
	g.currentBlock = g.nextBlock

	if g.blockFits() == false {
		g.GameOver = true
	}

	g.nextBlock = g.getRandomBlock()
	rowCleared := g.grid.ClearFullRows()

	if rowCleared > 0 {
		rl.PauseSound(g.clearSound)
		g.updateScore(rowCleared, 0)
	}
}

func (g *Game) blockFits() bool {
	titles := g.currentBlock.GetCellPostions()

	for _,item := range titles {
		if g.grid.IsCellEmpty(item.Row, item.Column) == false {
			return false
		}
	}
	
	return true
}

func (g *Game) reset() {
	g.grid.Initialize()
	g.blocks = g.getAllBlocks()
	g.currentBlock = g.getRandomBlock()
	g.nextBlock = g.getRandomBlock()
	g.Score = 0
}

func (g *Game) updateScore(linesCleared int, moveDownPoints int) {
	switch linesCleared {
		case 1 :
			g.Score += 100
		case 2:
			g.Score += 300
		case 3:
			g.Score += 500
		default:
			break
	}

	g.Score += moveDownPoints
}