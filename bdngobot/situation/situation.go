package situation

type Situation string

const (
	Any            Situation = "Any"
	ObstacleClose  Situation = "ObstacleClose"
	ObstacleMedium Situation = "ObstacleMedium"
	ObstacleFar    Situation = "ObstacleFar"
	MovingClose    Situation = "MovingClose"
	MovingMedium   Situation = "MovingMedium"
	MovingFar      Situation = "MovingFar"
)
