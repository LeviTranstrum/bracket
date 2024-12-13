//Main.cs

using BracketClasses;
using System.Collections.Generic;

public class Program
{
    public static void Main(string[] args)
    {


        var teams = new List<Team>
      {
        new Team("Arizona", "Diamondbacks", 3),
        new Team("Oakland", "Athletics", 4),
        new Team("Atlanta", "Braves", 2),
        new Team("Baltimore", "Orioles", 2),
        new Team("Boston", "Red Sox", 3),
        new Team("Chicago", "Cubs", 2),
        new Team("Chicago", "White Sox", 5),
        new Team("Cincinnati", "Reds", 4),
        new Team("Cleveland", "Guardians", 1),
        new Team("Colorado", "Rockies", 5),
        new Team("Detroit", "Tigers", 3),
        new Team("Houston", "Astros", 1),
        new Team("Kansas City", "Royals", 2),
        new Team("Los Angeles", "Angels", 5),
        new Team("Los Angeles", "Dodgers", 1),
        new Team("Miami", "Marlins", 5),
        new Team("Milwaukee", "Brewers", 1),
        new Team("Minnesota", "Twins", 4),
        new Team("New York", "Mets", 3),
        new Team("New York", "Yankees", 1),
        new Team("Philadelphia", "Phillies", 1),
        new Team("Pittsburgh", "Pirates", 5),
        new Team("San Diego", "Padres", 2),
        new Team("San Francisco", "Giants", 4),
        new Team("Seattle", "Mariners", 2),
        new Team("St. Louis", "Cardinals", 3),
        new Team("Tampa Bay", "Rays", 4),
        new Team("Texas", "Rangers", 3),
        new Team("Toronto", "Blue Jays", 5),
        new Team("Washington", "Nationals", 4)
       };

        Team.DisplayMatchTeams(teams);



    }
}