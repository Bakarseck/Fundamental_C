// See https://aka.ms/new-console-template for more information
using System.Collections.ObjectModel;
using System.Drawing;
using System.Runtime.InteropServices;

Random random = new Random();
int roll1 = random.Next(0, 10);

for (int i = 0; i < 10; i++)
{
    int roll = random.Next(0, 10);
    if (roll == roll1) {
        Console.BackgroundColor = ConsoleColor.Cyan;
    } else {
        Console.BackgroundColor = ConsoleColor.Red;
    }
    Console.WriteLine($"{roll}");
}
