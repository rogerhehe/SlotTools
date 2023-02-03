local config = 
{
	Drop = 
	{
		Empty = 0,
		Coin = 1,
		Blood = 2,
	},
	WeaponFire = 
	{
		FireType = 
		{
			Normal = 1,
			Random = 2,
		},
		FireNum = 
		{
			Single = 1,
			Multi = 2,
		},
	},
	BulletPath = 
	{
		Linear = 1,
		TargetPoint = 2,
		Tween = 3,
		Static = 4,
		Follow = 5,
	},
}
return config