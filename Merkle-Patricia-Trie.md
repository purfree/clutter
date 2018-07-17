# 梅克尔帕特里夏树（Merkle Patricia Trie）

节点类型如下
1. NULL (represented as the empty string)
2. branch A 17-item node [ v0 ... v15, vt ] 
3. leaf A 2-item node [ encodedPath, value ]
4. extension A 2-item node [ encodedPath, key ]

v0-v15对应0-15（16进制0-f），vt为存储的value

# 标记位
leaf和extension的结构一样，通过前缀标记位来区分，最低位0=偶数，1-奇数，低二位0=extension,1=leaf,占半个字节，所以如果是偶数，低四位填充0，如果是奇数，低四位填充key的第一个值。

    hex char    bits    |    node type partial     path length
    ----------------------------------------------------------
       0        0000    |       extension              even        
       1        0001    |       extension              odd         
       2        0010    |   terminating (leaf)         even        
       3        0011    |   terminating (leaf)         odd

# 示例
假设有4个key/value对，('do', 'verb'), ('dog', 'puppy'), ('doge', 'coin'), ('horse', 'stallion')。

首先将key转换为字节

    <64 6f> : 'verb'
    <64 6f 67> : 'puppy'
    <64 6f 67 65> : 'coin'
    <68 6f 72 73 65> : 'stallion'

数据库中的存储格式

    rootHash: [ <16>, hashA ]
    hashA:    [ <>, <>, <>, <>, hashB, <>, <>, <>, hashC, <>, <>, <>, <>, <>, <>, <>, <> ]
    hashC:    [ <20 6f 72 73 65>, 'stallion' ]
    hashB:    [ <00 6f>, hashD ]
    hashD:    [ <>, <>, <>, <>, <>, <>, hashE, <>, <>, <>, <>, <>, <>, <>, <>, <>, 'verb' ]
    hashE:    [ <17>, hashF ]
    hashF:    [ <>, <>, <>, <>, <>, <>, hashG, <>, <>, <>, <>, <>, <>, <>, <>, <>, 'puppy' ]
    hashG:    [ <35>, 'coin' ]

树结构如下

						                rootHash（extension，key=6长度1，奇数，因此标记位=1，6填充到低四位）
												  <16>
												    |
												    |
												  hashA（branch）
						0    1    2    3    4    5    6    7    8    9    a    b    c    d    e    f    '' （最后位值为空）
								   /			 \
								 /			   \
							       /			     \
							     /				       \
							   /					 \
							hashB					hashC
				         （extension，key=6f长度2，偶数		（leaf，key长度8，偶数
				           因此标记位=0，0填充到低四位）		    因此标记位=2,0填充到低四位）
							<00 6f>				<20 6f 72 73 65> 'stallion'
							   |
							   |
						   hashD（branch）
		0    1    2    3    4    5    6    7    8    9    a    b    c    d    e    f    'verb'
					      |
					      |
					hashE（extension）
					     <17>
					      |
			                      |
				            hashF（branch）
		0    1    2    3    4    5    6    7    8    9    a    b    c    d    e    f    'puppy'
					      |
					      |
					    hashG（leaf）
					  <35> 'coin'
