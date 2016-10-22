package situation

type Situation string

const (
	Any            Situation = "Any"
	ObstacleClose  Situation = "ObstacleClose"
	ObstacleMedium Situation = "ObstacleMedium"
	ObstacleFar    Situation = "ObstacleFar"
)
