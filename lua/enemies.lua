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
			Speed = 20.0,
			AttackSpeed = 0.5,
		},
		hp_per_wave = 5,
		damage_per_wave = 5,
		drop = 
		{
			coins = 
			{
				DropModel = 'Coin1',
				DropCount = 1,
				rate = 0.9,
			},
		},
		first_wave = 1,
	},
	[2] = 
	{
		name = 'Enemy2',
		character = 'Enemy2',
		attr = 
		{
			MaxHp = 80,
			Damage = 10,
			Armor = 0,
			Speed = 20.0,
			AttackSpeed = 0.5,
		},
		hp_per_wave = 5,
		damage_per_wave = 5,
		drop = 
		{
			coins = 
			{
				DropModel = 'Coin2',
				DropCount = 1,
				rate = 0.9,
			},
		},
		first_wave = 1,
	},
	[3] = 
	{
		name = 'Enemy3',
		character = 'Enemy3',
		attr = 
		{
			MaxHp = 120,
			Damage = 12,
			Armor = 0,
			Speed = 20.0,
			AttackSpeed = 0.5,
		},
		hp_per_wave = 5,
		damage_per_wave = 5,
		drop = 
		{
			coins = 
			{
				DropModel = 'Coin3',
				DropCount = 1,
				rate = 0.9,
			},
			blood = 
			{
				DropModel = 'BloodBag',
				DropCount = 0.1,
				rate = 50.0,
			},
		},
		first_wave = 1,
	},
	[4] = 
	{
		name = 'Enemy4',
		character = 'Enemy4',
		attr = 
		{
			MaxHp = 150,
			Damage = 18,
			Armor = 0,
			Speed = 20.0,
			AttackSpeed = 0.5,
		},
		hp_per_wave = 5,
		damage_per_wave = 5,
		drop = 
		{
			coins = 
			{
				DropModel = 'Coin4',
				DropCount = 1,
				rate = 0.9,
			},
		},
		first_wave = 1,
	},
	[5] = 
	{
		name = 'Enemy5',
		character = 'Enemy5',
		attr = 
		{
			MaxHp = 200,
			Damage = 20,
			Armor = 0,
			Speed = 20.0,
			AttackSpeed = 0.5,
		},
		hp_per_wave = 5,
		damage_per_wave = 5,
		drop = 
		{
			coins = 
			{
				DropModel = 'Coin5',
				DropCount = 1,
				rate = 0.9,
			},
		},
		first_wave = 1,
	},
	[6] = 
	{
		name = 'Enemy6',
		character = 'Enemy6',
		attr = 
		{
			MaxHp = 800,
			Damage = 30,
			Armor = 0,
			Speed = 20.0,
			AttackSpeed = 0.5,
		},
		hp_per_wave = 5,
		damage_per_wave = 5,
		drop = 
		{
			coins = 
			{
				DropModel = 'Coin6',
				DropCount = 1,
				rate = 0.9,
			},
		},
		first_wave = 1,
	},
}
return config