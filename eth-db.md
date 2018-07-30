    key														                    value
    LastBlock                                         最近的块hash，全节点的最近块hash。
                                                      官方解释 headBlockKey tracks the latest know full block's hash.
    LastHeader                                        最近的块hash，当前块hash，可能当前块还为加入区块链。
                                                      官方解释 headHeaderKey tracks the latest know header's hash.
    LastFast                                          最近的块hash，fast模式当前块hash，可能当前块还为加入区块链。
                                                      官方解释 headFastBlockKey tracks the latest known incomplete block's hash duirng fast sync.
    []byte("h") + number + []byte("n")                块号number的blockhash
    []byte("h") + number + blockhash                  块号number的header
    []byte("b") + number + blockhash                  块号number的body
    []byte("H") + blockhash                           块号number
    []byte("h") + number + blockhash + []byte("t")    到块号number的总难度，累加值
