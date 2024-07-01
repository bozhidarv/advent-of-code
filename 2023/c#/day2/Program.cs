internal class Program
{
    //#region day 2.1
    private const int MAX_RED_CUBES = 12;
    private const int MAX_GREEN_CUBES = 13;
    private const int MAX_BLUE_CUBES = 14;
    //#endregion

    private static void Main(string[] args)
    {
        var file = File.Open("../assets/day2.txt", FileMode.Open);
        var reader = new StreamReader(file);

        //#region day 2.1
        var validRoundIdSum = 0;
        //#endregion
        
        //#region day 2.2
        var sumOfPowerSets = 0;
        //#endregion
        
        while (reader.ReadLine() is { } line)
        {
            var lineSplitRaw = line.Split(":");
            var gameId = lineSplitRaw[0].Split("Game ")[1];
            var rounds = lineSplitRaw[1].Trim().Split(";");

            int maxRedCubes = 1, maxGreenCubes = 1, maxBlueCubes = 1;
            var validRound = true;
            foreach (var untrimmedRound in rounds)
            {
                var round = untrimmedRound.Trim();
                var cubes = round.Split(", ");
                foreach (var cube in cubes)
                {
                    var cubesNum = cube.Trim().Split(" ")[0];

                    //#region day 2.2
                    if (cube.Contains("red"))
                    {
                        maxRedCubes = Math.Max(maxRedCubes, int.Parse(cubesNum));
                    }
                    else if (cube.Contains("green"))
                    {
                        maxGreenCubes = Math.Max(maxGreenCubes, int.Parse(cubesNum));
                    }
                    else if (cube.Contains("blue"))
                    {
                        maxBlueCubes = Math.Max(maxBlueCubes, int.Parse(cubesNum));
                    }
                    //#endregion

                    //#region day 2.1
                    var roundValidityExp =
                        (cube.Contains("red") && int.Parse(cubesNum) > MAX_RED_CUBES) ||
                        (cube.Contains("green") && int.Parse(cubesNum) > MAX_GREEN_CUBES) ||
                        (cube.Contains("blue") && int.Parse(cubesNum) > MAX_BLUE_CUBES);

                    if (roundValidityExp)
                    {
                        validRound = false;
                    }
                    //#endregion
                }

            }

            //#region day 2.2
            sumOfPowerSets += maxRedCubes * maxGreenCubes * maxBlueCubes;
            //#endregion

            //#region day 2.1
            if (validRound) validRoundIdSum += int.Parse(gameId);
            //#endregion
        }

        Console.WriteLine(sumOfPowerSets);
    }
}
