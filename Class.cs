//Class.cs

using System;
using System.Collections.Generic;

namespace BracketClasses
{
    public class Team
    {
        public string City { get; set; }
        public string TeamName { get; set; }
        public int Ranking { get; set; }

        public Team(string city, string teamName, int ranking)
        {
            City = city;
            TeamName = teamName;
            Ranking = ranking;
        }

        public static void DisplayMatchTeams(List<Team> teams)
        {
            foreach (var team in teams)
            {
                Console.WriteLine($"The {team.City} {team.TeamName} have a ranking of {team.Ranking}.");
            }
        }
        public static void TeamBattle(List<Team> teams)
        {
            Random random = new();
            int Team1 = random.Next(teams.Count);
            int Team2 = random.Next(teams.Count);


            if (teams[Team1].Ranking > teams[Team2].Ranking)
            {
                Console.WriteLine($"The {teams[Team1].City} {teams[Team1].TeamName} have a ranking of {teams[Team1].Ranking}.");
                Console.WriteLine($"The {teams[Team2].City} {teams[Team2].TeamName} have a ranking of {teams[Team2].Ranking}.");
                Console.WriteLine($"The {teams[Team1].City} {teams[Team1].TeamName} are higher ranked.");
            }

            else
            {
                Console.WriteLine($"The {teams[Team2].City} {teams[Team2].TeamName} have a ranking of {teams[Team2].Ranking}.");
                Console.WriteLine($"The {teams[Team1].City} {teams[Team1].TeamName} have a ranking of {teams[Team1].Ranking}.");
                Console.WriteLine($"The {teams[Team2].City} {teams[Team2].TeamName} are higher ranked.");
            }

        }

    }
}

