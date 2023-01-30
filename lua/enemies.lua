local config = 
{
	[1] = 
	{
		name = 'Enemy1',
		character = 'Enemy1',
		attr = 
		{
			MaxHp = 50,
			Damage = 10,
			Armor = 0,
			Speed = 10.0,
			AttackSpeed = 0.5,
		},
		hp_per_wave = 5,
		damage_per_wave = 5,
		drop = 
		{
			Coin = 
			{
				DropModel = 'Coin1',
				DropCount = 1,
				Rate = 0.3,
			},
			Blood = 
			{
				DropModel = 'BloodBag',
				DropCount = 50.0,
				Rate = 0.1,
			},
		},
		first_wave = 1,
	},
	[2] = 
	{
		name = 'Enemy2',
		character = 'Enemy1',
		attr = 
		{
			MaxHp = 80,
			Damage = 10,
			Armor = 0,
			Speed = 10.0,
			AttackSpeed = 0.5,
		},
		hp_per_wave = 5,
		damage_per_wave = 5,
		drop = 
		{
			Coin = 
			{
				DropModel = 'Coin2',
				DropCount = 1,
				Rate = 0.2,
			},
		},
		first_wave = 1,
	},
	[3] = 
	{
		name = 'Enemy3',
		character = 'Enemy1',
		attr = 
		{
			MaxHp = 120,
			Damage = 12,
			Armor = 0,
			Speed = 10.0,
			AttackSpeed = 0.5,
		},
		hp_per_wave = 5,
		damage_per_wave = 5,
		drop = 
		{
			Coin = 
			{
				DropModel = 'Coin3',
				DropCount = 1,
				Rate = 0.4,
			},
			Blood = 
			{
				DropModel = 'BloodBag',
				DropCount = 50.0,
				Rate = 0.1,
			},
		},
		first_wave = 1,
	},
	[4] = 
	{
		name = 'Enemy4',
		character = 'Enemy1',
		attr = 
		{
			MaxHp = 150,
			Damage = 18,
			Armor = 0,
			Speed = 10.0,
			AttackSpeed = 0.5,
		},
		hp_per_wave = 5,
		damage_per_wave = 5,
		drop = 
		{
			Coin = 
			{
				DropModel = 'Coin4',
				DropCount = 1,
				Rate = 0.4,
			},
		},
		first_wave = 1,
	},
	[5] = 
	{
		name = 'Enemy5',
		character = 'Enemy1',
		attr = 
		{
			MaxHp = 200,
			Damage = 20,
			Armor = 0,
			Speed = 10.0,
			AttackSpeed = 0.5,
		},
		hp_per_wave = 5,
		damage_per_wave = 5,
		drop = 
		{
			Coin = 
			{
				DropModel = 'Coin5',
				DropCount = 1,
				Rate = 0.5,
			},
		},
		first_wave = 1,
	},
	[6] = 
	{
		name = 'EnemyBoss1',
		character = 'Enemy1',
		attr = 
		{
			MaxHp = 800,
			Damage = 30,
			Armor = 0,
			Speed = 10.0,
			AttackSpeed = 0.5,
		},
		hp_per_wave = 5,
		damage_per_wave = 5,
		drop = 
		{
			Coin = 
			{
				DropModel = 'Coin6',
				DropCount = 1,
				Rate = 0.2,
			},
		},
		first_wave = 1,
	},
}
return config