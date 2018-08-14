package magic_cube

// 第二种思路，按真实的块分，一共3*9=27块，中间一块看不到
// 一共有三维坐标，中间的看不到的块为（0，0，0）, 右手象系
type Block interface {
	// 一块
}

type BlockCell struct {
	Block

	X int
	Y int
	Z int
}

type BlockCellZero struct {
	// 中心的，（0，0，0）
	BlockCell
}

func (b *BlockCellZero) Init() {
	b.X = 0
	b.Y = 0
	b.Z = 0
}

type BlockCellCenter struct {
	// 每一个面的中心，三个坐标，其中两个是0，
	// 只有一个color
	BlockCell
	color int
}

type BlockCellEdge struct {
	// 棱
	// 两个color
	// 其中一个坐标是0
}

type BlockCellApex struct {
	// 点，
	// 三个color
	// 三个坐标都不为零
}

type MagicCube2 struct {
}
