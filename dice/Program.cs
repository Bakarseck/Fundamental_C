using System;
using System.Windows.Forms;

namespace DiceRollGame
{
    public partial class Form1 : Form
    {
        private Random random = new Random();

        public Form1()
        {
            InitializeComponent();
        }

        private void RollButton_Click(object sender, EventArgs e)
        {
            // Générer un nombre aléatoire entre 1 et 6 (inclus)
            int result = random.Next(1, 7);

            // Afficher le résultat
            resultLabel.Text = "Résultat : " + result.ToString();
        }
    }
}

